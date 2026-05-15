# EarthquakeCatalog SDK utility: make_error

from __future__ import annotations
from core.operation import EarthquakeCatalogOperation
from core.result import EarthquakeCatalogResult
from core.control import EarthquakeCatalogControl
from core.error import EarthquakeCatalogError


def make_error_util(ctx, err):
    if ctx is None:
        from core.context import EarthquakeCatalogContext
        ctx = EarthquakeCatalogContext({}, None)

    op = ctx.op
    if op is None:
        op = EarthquakeCatalogOperation({})
    opname = op.name
    if opname == "" or opname == "_":
        opname = "unknown operation"

    result = ctx.result
    if result is None:
        result = EarthquakeCatalogResult({})
    result.ok = False

    if err is None:
        err = result.err
    if err is None:
        err = ctx.make_error("unknown", "unknown error")

    errmsg = ""
    if isinstance(err, EarthquakeCatalogError):
        errmsg = err.msg
    elif hasattr(err, "msg") and err.msg is not None:
        errmsg = err.msg
    elif isinstance(err, str):
        errmsg = err
    else:
        errmsg = str(err)

    msg = "EarthquakeCatalogSDK: " + opname + ": " + errmsg
    msg = ctx.utility.clean(ctx, msg)

    result.err = None

    spec = ctx.spec

    if ctx.ctrl.explain is not None:
        ctx.ctrl.explain["err"] = {"message": msg}

    sdk_err = EarthquakeCatalogError("", msg, ctx)
    sdk_err.result = ctx.utility.clean(ctx, result)
    sdk_err.spec = ctx.utility.clean(ctx, spec)

    if isinstance(err, EarthquakeCatalogError):
        sdk_err.code = err.code

    ctx.ctrl.err = sdk_err

    if ctx.ctrl.throw_err is False:
        return result.resdata, None

    return None, sdk_err
